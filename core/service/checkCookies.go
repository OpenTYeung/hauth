package service

import (
	"net/http"

	"github.com/asofdate/hauth/utils/jwt"
)

const redirect = `
<script type="text/javascript">
    $.Hconfirm({
		cancelBtn:false,
        header:"连接已断开",
        body:"用户连接已断开，请重新登录",
        callback:function () {
            window.location.href="/"
        }
    })
</script>
`

func CheckConnection(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Authorization")
	if err != nil || !jwt.CheckToken(cookie.Value) {
		w.Write([]byte(redirect))
	}
}
