#!/bin/bash

get_matrix() {
  GROUP=$1
  NAME=$2
  wget https://suitesparse-collection-website.herokuapp.com/MM/$GROUP/$NAME.tar.gz
  tar xzf $NAME.tar.gz
  mv $NAME/$NAME.mtx .
  rm -rf $NAME
  rm $NAME.tar.gz
}

# M1
# get_matrix GHS_psdef inline_1

# M3
# get_matrix GHS_psdef audikw_1

# M5
# get_matrix Janna Flan_1565

# M8
# get_matrix Bourchtein atmosmodl 

# complex
#get_matrix Chevron Chevron4

# unsymmetric
get_matrix CEMW tmt_unsym