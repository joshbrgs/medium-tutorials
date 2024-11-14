#!/bin/bash

# List of project directories
projects=("./host" "./navigator" "./dashboard")

# Loop through each directory
for project in "${projects[@]}"; do
  # Navigate to the project directory
  echo "Navigating to $project"
  cd "$project" || { echo "Failed to enter $project"; continue; }

  # Check if package.json exists to confirm it's a Node.js project
  if [ -f "package.json" ]; then
    echo "Installing new packages in $project"
    npm i
    echo "Starting project in $project"
    npm run start &
  else
    echo "No package.json found in $project. Skipping..."
  fi

  # Go back to the original directory
  cd - > /dev/null || exit
done

# Wait for all background processes to finish
wait
echo "All projects started."

