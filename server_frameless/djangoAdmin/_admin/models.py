from django.db import models
from django.contrib.gis.db import models as gis_models

class City(models.Model):
    name = models.TextField(primary_key=True)

    class Meta:
        db_table = 'vetrof.cities'
        verbose_name = 'City'
        verbose_name_plural = 'Cities'

class Place(models.Model):
    id = models.AutoField(primary_key=True)
    name = models.TextField(null=True, blank=True)
    description = models.TextField(null=True, blank=True)


    desc = models.TextField(null=True, blank=True)
    distance = models.DecimalField(max_digits=10, decimal_places=2, null=True, blank=True)
    created_at = models.DateTimeField(null=True, blank=True)
    updated_at = models.DateTimeField(null=True, blank=True)
    deleted_at = models.DateTimeField(null=True, blank=True)
    city_name = models.ForeignKey(City, on_delete=models.SET_NULL, null=True, blank=True, db_column='city_name')

    class Meta:
        db_table = 'vetrof.places'
        verbose_name = 'Place'
        verbose_name_plural = 'Places'
