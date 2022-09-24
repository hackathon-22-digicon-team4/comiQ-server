{{- define "query" -}}
{{- $table := . -}}
func Iterate{{ $table.GoName }}(sc interface{ Scan(...interface{}) error}) ({{ $table.GoName }}, error) {
  t := {{ $table.GoName }}{}
  if err := sc.Scan(t.Ptrs()...); err != nil {
    return {{ $table.GoName }}{}, MapError(err)
  }
  return t, nil
}

{{ range $index := $table.Indexes -}}
func Select{{ if $index.IsUnique }}One{{ end }}{{ $table.GoName }}By{{ range $i, $f := $index.Fields }} {{- $f.GoName -}} {{- if not ($index.Tail $i) -}} And {{- end }}{{ end }}(ctx context.Context, txn *sql.Tx, {{ range $i, $f := $index.Fields }} {{- $f.Name }} {{ $f.GoType }} {{- if not ($index.Tail $i) -}} , {{- end }}{{ end }}) ({{ if not $index.IsUnique -}} []* {{- end }} {{- $table.GoName }}, error) {
  query, params, err := squirrel.
    Select({{ $table.GoName }}AllColumns...).
    From({{ $table.GoName}}TableName).
    Where(squirrel.Eq{
      {{- range $index.Fields }}
      "{{ .Name }}": {{ .Name }},
      {{- end }}
    }).
    ToSql()
  if err != nil {
    return {{ if not $index.IsUnique -}} nil {{- else -}} {{ $table.GoName }}{} {{- end -}}, MapError(err)
  }
  stmt, err := txn.PrepareContext(ctx, query)
	if err != nil {
    return {{ if not $index.IsUnique -}} nil {{- else -}} {{ $table.GoName }}{} {{- end -}}, MapError(err)
	}

  {{- if $index.IsUnique }}
  return Iterate{{ $table.GoName }}(stmt.QueryRowContext(ctx, params...))
  {{ else }}
  rows, err := stmt.QueryContext(ctx, params...)
  if err != nil {
    return nil, MapError(err)
  }
  res := make([]*{{ $table.GoName }}, 0)
  for rows.Next() {
    t, err := Iterate{{ $table.GoName }}(rows)
    if err != nil {
      return nil, MapError(err)
    }
    res = append(res, &t)
  }
  return res, nil
  {{ end }}
}

{{ end }}
{{ end }}