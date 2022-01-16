module bbtest

go 1.15

require (
    "bbtest/storage" v0.0.0
    "bbtest/script"  v0.0.0
    "bbtest/scripttool" v0.0.0
)

replace (
    "bbtest/storage" v0.0.0 => "./storage"
    "bbtest/script"  v0.0.0 => "./script"
    "bbtest/scripttool" v0.0.0 => "./scripttool"
)