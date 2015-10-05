#!/usr/bin/env bats

@test "Check for available command" {
    command -v takegur
}

@test "black twitter test." {
    run bash -c "takegur black-twitter"
    [ $status = 0 ]
}

@test "mrw test." {
    run bash -c "takegur mrw"
    [ $status = 0 ]
}

@test "epic fails test." {
    run bash -c "takegur fails"
    [ $status = 0 ]
}

@test "wallpaper test." {
    run bash -c "takegur wallpapers"
    [ $status = 0 ]
}

@test "sc test." {
    run bash -c "takegur stay-classy"
    [ $status = 0 ]
}

@test "darwin test." {
    run bash -c "takegur darwin-awards"
    [ $status = 0 ]
}

@test "dickbutt test." {
    run bash -c "takegur dickbutt"
    [ $status = 0 ]
}