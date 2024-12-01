"""Add Bookmark and Tag models

Revision ID: 1ee2f0d8936c
Revises: 1a31ce608336
Create Date: 2024-12-01 16:47:04.281146

"""
from alembic import op
import sqlalchemy as sa
import sqlmodel.sql.sqltypes


# revision identifiers, used by Alembic.
revision = '1ee2f0d8936c'
down_revision = '1a31ce608336'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table('bookmark',
    sa.Column('id', sa.Uuid(), nullable=False),
    sa.Column('url', sqlmodel.sql.sqltypes.AutoString(length=1000), nullable=False),
    sa.Column('title', sqlmodel.sql.sqltypes.AutoString(length=50), nullable=True),
    sa.Column('description', sqlmodel.sql.sqltypes.AutoString(length=1000), nullable=True),
    sa.Column('owner_id', sa.Uuid(), nullable=False),
    sa.ForeignKeyConstraint(['owner_id'], ['user.id'], ondelete='CASCADE'),
    sa.PrimaryKeyConstraint('id')
    )
    op.create_table('tag',
    sa.Column('id', sa.Uuid(), nullable=False),
    sa.Column('name', sqlmodel.sql.sqltypes.AutoString(length=255), nullable=False),
    sa.Column('owner_id', sa.Uuid(), nullable=False),
    sa.Column('bookmark_id', sa.Uuid(), nullable=False),
    sa.ForeignKeyConstraint(['bookmark_id'], ['bookmark.id'], ondelete='CASCADE'),
    sa.ForeignKeyConstraint(['owner_id'], ['user.id'], ondelete='CASCADE'),
    sa.PrimaryKeyConstraint('id')
    )
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_table('tag')
    op.drop_table('bookmark')
    # ### end Alembic commands ###
