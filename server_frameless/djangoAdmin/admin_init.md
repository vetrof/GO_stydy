python manage.py inspectdb > _admin/models.py  
python manage.py makemigrations _admin  
python manage.py migrate _admin --fake-initial 
python manage.py migrate
