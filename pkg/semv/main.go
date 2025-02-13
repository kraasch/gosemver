
package semv

import (
  // "fmt"
)

const (
  url = "https://go.dev/doc/devel/release"
)

type VersionData = struct{
  Number   string
  Semver   string
  Gosemver string
  Date     string
}

var Versions = []VersionData{
  {"25","1.23", "go1.23.0", "2024-08-13"},
  {"24","1.22", "go1.22.0", "2024-02-06"},
  {"23","1.21", "go1.21.0", "2023-08-08"},
  {"22","1.20", "go1.20",   "2023-02-01"},
  {"21","1.19", "go1.19",   "2022-08-02"},
  {"20","1.18", "go1.18",   "2022-03-15"},
  {"19","1.17", "go1.17",   "2021-08-16"},
  {"18","1.16", "go1.16",   "2021-02-16"},
  {"17","1.15", "go1.15",   "2020-08-11"},
  {"16","1.14", "go1.14",   "2020-02-25"},
  {"15","1.13", "go1.13",   "2019-09-03"},
  {"14","1.12", "go1.12",   "2019-02-25"},
  {"13","1.11", "go1.11",   "2018-08-24"},
  {"12","1.10", "go1.10",   "2018-02-16"},
  {"11", "1.9", "go1.9",    "2017-08-24"},
  {"10", "1.8", "go1.8",    "2017-02-16"},
  { "9", "1.7", "go1.7",    "2016-08-15"},
  { "8", "1.6", "go1.6",    "2016-02-17"},
  { "7", "1.5", "go1.5",    "2015-08-19"},
  { "6", "1.4", "go1.4",    "2014-12-10"},
  { "5", "1.3", "go1.3",    "2014-06-18"},
  { "4", "1.2", "go1.2",    "2013-12-01"},
  { "3", "1.1", "go1.1",    "2013-05-13"},
  { "2", "1.0", "go1",      "2012-03-28"},
  { "1", "0.0", "pre.go1",  "2011-03-16"},
}

func SemverToUrl(semver string) string {
  gosemver := ""
  for _, v := range Versions {
    if v.Semver == semver {
      gosemver = v.Gosemver
    }
  }
  if gosemver != "" {
    return url + "#" + gosemver
  } else {
    return url
  }
}

func DateToSemver(date string) string {
  semver := ""
  for _, v := range Versions {
    semver = v.Semver
    if v.Date <= date {
      break
    }
  }
  return semver
}

