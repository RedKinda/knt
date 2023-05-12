# Generated by Django 4.2.1 on 2023-05-12 13:32

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("knt_backend", "0005_transactionitem"),
    ]

    operations = [
        migrations.AlterField(
            model_name="product",
            name="id",
            field=models.AutoField(editable=False, primary_key=True, serialize=False),
        ),
        migrations.AlterField(
            model_name="receipt",
            name="id",
            field=models.AutoField(editable=False, primary_key=True, serialize=False),
        ),
        migrations.AlterField(
            model_name="receipt",
            name="timestamp",
            field=models.DateTimeField(auto_now_add=True),
        ),
        migrations.AlterField(
            model_name="taxcategory",
            name="id",
            field=models.AutoField(editable=False, primary_key=True, serialize=False),
        ),
        migrations.AlterField(
            model_name="transaction",
            name="id",
            field=models.AutoField(editable=False, primary_key=True, serialize=False),
        ),
        migrations.AlterField(
            model_name="user",
            name="id",
            field=models.AutoField(editable=False, primary_key=True, serialize=False),
        ),
    ]
