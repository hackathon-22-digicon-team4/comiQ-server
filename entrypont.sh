# Secret Managerから文字列で渡ってくる環境変数jsonをパースして、環境変数にいれる
if [ -n "$JSON_FROM_SECRET_MANAGER_STR" ]
then
    echo ${JSON_FROM_SECRET_MANAGER_STR} | jq -r 'to_entries|map("\(.key)=\(.value|tostring)")|.[]' > /tmp/secrets.env
    eval $(cat /tmp/secrets.env | sed 's/^/export /')
    rm -f /tmp/secrets.env
fi

