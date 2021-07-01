function not_aligned() {
  echo -e "\nHmmm, it looks the fields of certain structs can be re-ordered\n"
  echo -e "You can use the \`fieldalignment\` tool to apply the recommended reordering as follows:\n"
  echo -e "\tUsage: fieldalignment -fix /path/to/file/having/struct"
}

function command_not_found() {
  echo -e "The command fieldalignment doesn't seem to be installed.\n"
  echo -e "\tIntsallation: go get golang.org/x/tools/go/analysis/passes/fieldalignment\n"
}

if ! type fieldalignment &> /dev/null; then
  command_not_found
  exit 1
fi

if ! fieldalignment ./... ; then
  not_aligned
  exit 1
fi