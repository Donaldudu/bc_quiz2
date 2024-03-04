# bc_quiz2
echo "# bc_quiz2" >> README.md
git init
git add *
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Donaldudu/bc_quiz2.git
git push -u origin main
git branch
git branch dev
git checkout dev
git add *
git commit -m "second commit"
git push -u origin dev
git checkout main
git merge dev
git commit -m "dev and main merged"
git push origin main
