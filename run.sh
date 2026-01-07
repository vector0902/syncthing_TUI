
export SYNCTHING_API_KEY="bJJfpeAN4WVPqTsePxvkdcMv59Mv2bPt"
export SYNCTHING_URL="http://localhost:8666"

export SYNCTHING_API_KEY="ccx3n5s64Et73nGC2ouREcSXafDG6JNo"
export SYNCTHING_URL="http://localhost:8388"

[ -e syncthing_TUI ] || go build

./syncthing_TUI

