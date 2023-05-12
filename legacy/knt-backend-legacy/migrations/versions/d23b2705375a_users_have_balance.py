"""users have balance

Revision ID: d23b2705375a
Revises: e2fe825de131
Create Date: 2022-01-07 21:26:57.257628

"""
from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision = 'd23b2705375a'
down_revision = 'e2fe825de131'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('user', sa.Column('balance', sa.Numeric(), nullable=True))
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('user', 'balance')
    # ### end Alembic commands ###