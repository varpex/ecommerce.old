# ~/.profile: executed by Bourne-compatible login shells.

if [ "$BASH" ]; then
  if [ -f ~/.bashrc ]; then
    . ~/.bashrc
  fi
fi

mesg n || true

# export PATH="$PATH:$GOPATH/bin"

# export GOROOT=/usr/local/go

# export PATH=/home/bin:$GOROOT/bin:$PATH

export PATH=$PATH:/usr/local/go/bin

# export GOPATH=/home/golang
# export PATH=\$PATH:\$GOPATH/bin

# add follows to the end (set proxy settings to the environment variables)
MY_PROXY_URL="http://iruser717439:1369s1r3d691369@us.mybestport.com:443/"
HTTP_PROXY=$MY_PROXY_URL
HTTPS_PROXY=$MY_PROXY_URL
FTP_PROXY=$MY_PROXY_URL
http_proxy=$MY_PROXY_URL
https_proxy=$MY_PROXY_URL
ftp_proxy=$MY_PROXY_URL
export HTTP_PROXY HTTPS_PROXY FTP_PROXY http_proxy https_proxy ftp_proxy
