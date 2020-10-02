#!/bin/bash

get_matrix() {
  GROUP=$1
  NAME=$2
  wget https://suitesparse-collection-website.herokuapp.com/MM/$GROUP/$NAME.tar.gz
  tar xzf $NAME.tar.gz
  mv $NAME/$NAME.mtx .
  rmdir $NAME
  rm $NAME.tar.gz
}

get_matrix GHS_psdef inline_1