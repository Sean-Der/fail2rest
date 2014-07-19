#!/bin/sh
# kFreeBSD do not accept scripts as interpreters, using #!/bin/sh and sourcing.
if [ true != "$INIT_D_SCRIPT_SOURCED" ] ; then
    set "$0" "$@"; INIT_D_SCRIPT_SOURCED=true . /lib/init/init-d-script
fi
### BEGIN INIT INFO
# Provides:          fail2rest
# Required-Start:    $remote_fs $syslog
# Required-Stop:     $remote_fs $syslog
# Should-Start:	     fail2ban
# Should-Stop:       fail2ban
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: fail2rest initscript
# Description:       fail2rest is a small
#                    REST server that aims
#                    to allow full administration
#                    of a fail2ban server via HTTP
#
### END INIT INFO


USER="root"
#FIXME
GOPATH="GOPATH"
WORKDIR="/var/run/fail2ban"
#FIXME path to your fail2rest binary
DAEMON="FAIL2REST_BINARY"
CONFIG="/etc/fail2rest.json"

# Author: Sean DuBois <sean@siobud.com>
#
DESC="fail2ban REST server"
NAME="fail2rest"

case "$1" in
      start)
        echo "Starting $NAME ..."
        if [ -f "$WORKDIR/$NAME.pid" ]
        then
            echo "Already running according to $WORKDIR/$NAME.pid"
            exit 1
        fi
        cd "$WORKDIR"
  export GOPATH="$GOPATH"
  export PATH="PATH=/usr/sbin:/usr/bin:/sbin:/bin:$GOPATH/bin"
        /bin/su -m -l $USER -c "$DAEMON --config $CONFIG" > "$WORKDIR/$NAME.log" 2>&1 &
        PID=$!
        echo $PID > "$WORKDIR/$NAME.pid"
        echo "Started with pid $PID - Logging to $WORKDIR/$NAME.log" && exit 0
        ;;
      stop)
        echo "Stopping $NAME ..."
        if [ ! -f "$WORKDIR/$NAME.pid" ]
        then
            echo "Already stopped!"
            exit 1
        fi
        PID=`cat "$WORKDIR/$NAME.pid"`
        kill $PID
        rm -f "$WORKDIR/$NAME.pid"
        echo "stopped $NAME" && exit 0
        ;;
      restart)
        $0 stop
        sleep 1
        $0 start
        ;;
      status)
        if [ -f "$WORKDIR/$NAME.pid" ]
        then
            PID=`cat "$WORKDIR/$NAME.pid"`
            if [ "$(/bin/ps --no-headers -p $PID)" ]
            then
                echo "$NAME is running (pid : $PID)" && exit 0
            else
                echo "Pid $PID found in $WORKDIR/$NAME.pid, but not running." && exit 1
            fi
        else
            echo "$NAME is NOT running" && exit 1
        fi
    ;;
      *)
      echo "Usage: /etc/init.d/$NAME {start|stop|restart|status}" && exit 1
      ;;
esac

exit 0
