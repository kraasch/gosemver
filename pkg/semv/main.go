
package semv

import (
  // "fmt"
)

const (
  url = "https://go.dev/doc/devel/release"
)

var (
  // all versions.
  versions = []struct{
    semver   string
    gosemver string
    date     string
  }{
    {"1.23", "go1.23.0", "2024-08-13"},
    {"1.22", "go1.22.0", "2024-02-06"},
    {"1.21", "go1.21.0", "2023-08-08"},
    {"1.20", "go1.20",   "2023-02-01"},
    {"1.19", "go1.19",   "2022-08-02"},
    {"1.18", "go1.18",   "2022-03-15"},
    {"1.17", "go1.17",   "2021-08-16"},
    {"1.16", "go1.16",   "2021-02-16"},
    {"1.15", "go1.15",   "2020-08-11"},
    {"1.14", "go1.14",   "2020-02-25"},
    {"1.13", "go1.13",   "2019-09-03"},
    {"1.12", "go1.12",   "2019-02-25"},
    {"1.11", "go1.11",   "2018-08-24"},
    {"1.10", "go1.10",   "2018-02-16"},
    { "1.9", "go1.9",    "2017-08-24"},
    { "1.8", "go1.8",    "2017-02-16"},
    { "1.7", "go1.7",    "2016-08-15"},
    { "1.6", "go1.6",    "2016-02-17"},
    { "1.5", "go1.5",    "2015-08-19"},
    { "1.4", "go1.4",    "2014-12-10"},
    { "1.3", "go1.3",    "2014-06-18"},
    { "1.2", "go1.2",    "2013-12-01"},
    { "1.1", "go1.1",    "2013-05-13"},
    { "1.0", "go1",      "2012-03-28"},
    { "0.0", "pre.go1",  "2011-03-16"},
  }
)

func SemverToUrl(semver string) string {
  gosemver := ""
  for _, v := range versions {
    if v.semver == semver {
      gosemver = v.gosemver
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
  for _, v := range versions {
    semver = v.semver
    if v.date <= date {
      break
    }
  }
  return semver
}

