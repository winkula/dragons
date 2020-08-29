<template>
  <div>
    <Header></Header>
    <Menu @new="newGame" @difficulty="changeDifficulty" @help="showHelp"></Menu>
    <Grid :grid="game.puzzle" @filled="fillCell" :is-valid="isValid" :is-solved="isSolved"></Grid>
    <CellSelect v-model="fillType"></CellSelect>
    <Footer :copyright="copyright"></Footer>
    <Help :visible="helpVisible"></Help>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

import Header from "./Header.vue";
import Grid from "./Grid.vue";
import Menu from "./Menu.vue";
import CellSelect from "./CellSelect.vue";
import Footer from "./Footer.vue";
import Help from "./Help.vue";

import {
  createGame,
  Cell,
  CellDefinition,
  CellType,
  getCellType,
} from "../logic";
import { GameStatus } from "../logic/game";

export default Vue.extend({
  components: {
    Header,
    Menu,
    Grid,
    CellSelect,
    Footer,
    Help,
  },
  data() {
    return {
      game: null,
      isValid: true,
      isSolved: false,
      fillType: getCellType(CellType.Empty).value,
      helpVisible: false,
      copyright: "© 2020 Mathias Winkler",
    };
  },
  methods: {
    fillCell(cell: Cell) {
      if (cell.given) {
        // cannot change a cell that is given
        return;
      }

      if (this.fillType == null) {
        // cannot set a cell, when no fill type is selected
        return;
      }

      if (cell.isDefined) {
        // set back to undeinef
        cell.value = getCellType(CellType.Undefined).value;
      } else {
        // set value of cell
        cell.value = this.fillType;
      }

      // validate game
      const status = this.game.status;

      this.isValid = status !== GameStatus.Invalid;
      this.isSolved = status === GameStatus.Solved;
    },
    newGame() {
      this.game = createGame();
      this.isValid = true;
      this.isSolved = false;
    },
    changeDifficulty() {},
    showHelp() {
      this.helpVisible = !this.helpVisible;
    },
  },
  created() {
    this.newGame();
  }  
});
</script>