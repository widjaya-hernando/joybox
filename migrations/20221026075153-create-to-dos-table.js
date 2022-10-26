'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    return queryInterface.createTable('to_dos', {
      id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
        primaryKey: true,
        autoIncrement: true,
      },
      user_id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
      },
      name: {
          type: Sequelize.STRING,
          allowNull: false,
      },
      description:{
          type: Sequelize.TEXT,
          allowNull: true,
      },
      due_date:{
        type: Sequelize.DATE,
        allowNull: false,
      },
      status:{
        type: Sequelize.BOOLEAN,
        allowNull: false,
      }
    })
  },

  async down (queryInterface, Sequelize) {
    return queryInterface.dropTable('to_do')
  }
};
