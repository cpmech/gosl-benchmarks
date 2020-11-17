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

# unsymmetric 1
#get_matrix CEMW tmt_unsym

# unsymmetric 2
#get_matrix Hamrle Hamrle3

# unsymmetric 3
#get_matrix ATandT pre2

# unsymmetric 4
get_matrix Janna ML_Laplace
