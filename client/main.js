import WebFont from 'webfontloader';
import loglevel from 'loglevel';

import './loginit';
import version from './version';

import Manager from './manager';
import Game from './game';

const logger = loglevel.getLogger('main');

/*
// https://learn.javascript.ru/cookie
// возвращает cookie с именем name или undefined
function getCookie(name) {
  const matches = document.cookie.match(
    new RegExp(`(?:^|; )${name.replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, '\\$1')}=([^;]*)`),
  );
  return matches ? decodeURIComponent(matches[1]) : undefined;
}
const session = getCookie('session');
*/

logger.info(
  `Unbubblable v. ${version.version} (build ${version.build}) released on ${version.date}`,
);

const start = function start() {
  const game = new Game();
  game.animate();
};

const fail = function fail() {
  // Браузер не соответствует требованиям
  // const container = document.getElementById('container');
  // FIX: log.domElement может измениться
  // container.appendChild(log.box);
};

const manager = new Manager(start, fail);
if (!manager.require()) {
  fail();
  throw new Error('Браузер не соответствует требованиям');
}

// FIX: Странная конструкция
window.onload = manager.windowload.bind(manager);

WebFont.load({
  custom: {
    families: ['Hobo'],
  },
  classes: false,
  active() {
    manager.fontload(true);
  },
  inactive() {
    manager.fontload(false);
  },
});
