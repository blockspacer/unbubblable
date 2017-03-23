class Tile {

    constructor(ground, block, detail) {

        this.ground = ground || 0;
        this.block = block || 0;
        this.detail = detail || 0;

    }

    set(ground, block, detail) {

        this.ground = ground;
        this.block = block;
        this.detail = detail;

        return this;

    }

    copy(tile) {

        this.ground = tile.ground;
        this.block = tile.block;
        this.detail = tile.detail;

        return this;

    }

    clone() {

        return new Tile(this.ground, this.block, this.detail);

    }

};

class Chunk {

    constructor(terrain, msgChunk) {

        this.terrain = terrain;

        let chunkedWidth = Math.ceil(terrain.width / terrain.chunkSize);

        let y = Math.floor(msgChunk.index / chunkedWidth)
        let x = msgChunk.index - (y * chunkedWidth)

        this.meshX = x * terrain.chunkSize;
        this.meshY = y * terrain.chunkSize;


        let size = terrain.chunkSize * terrain.chunkSize;
        this.tiles = new Array(size);
        for (let i = 0; i < size; i++) {
            this.tiles[i] = new Tile(msgChunk.tiles[i].ground, msgChunk.tiles[i].block, msgChunk.tiles[i].detail);
        }

        this.calcMesh();

    }

    calcMesh() {

        // Ландшафт
        let geoLandscape = new THREE.Geometry();
        geoLandscape.vertices = this.terrain.verticesBuffer;

        // Детали
        let geoDetails = new THREE.Geometry();

        // quad добавляет в геометрию квадрат по 4 индексам вершин
        function quad(geo, atlas, tile, v0, v1, v2, v3) {

            // Создаем 2 треугольника
            geo.faces.push(new THREE.Face3(v0, v2, v1));
            geo.faces.push(new THREE.Face3(v2, v3, v1));
            // Задаем текстурные координаты
            geo.faceVertexUvs[0].push(atlas.tiles[tile].faces[0][0]);
            geo.faceVertexUvs[0].push(atlas.tiles[tile].faces[0][1]);

        }

        // detail добавляет в геометрию деталь (два перпендикулярных квадрата)
        function detail(geo, atlas, tile, x, y) {

            let geoDec = new THREE.Geometry();
            // Заполняем вершинный буфер
            geoDec.vertices.push(new THREE.Vector3(-0.5, 0, 1));
            geoDec.vertices.push(new THREE.Vector3(0.5, 0, 1));
            geoDec.vertices.push(new THREE.Vector3(-0.5, 0, 0));
            geoDec.vertices.push(new THREE.Vector3(0.5, 0, 0));

            // Создаем первую плоскость
            quad(geoDec, atlas, tile, 0, 1, 2, 3);

            // Добавляем вторую плоскость, перпендикулярную первой
            geoDec.merge(geoDec.clone(), new THREE.Matrix4().makeRotationZ(Math.PI / 2));

            // Перемещаем
            let mt = new THREE.Matrix4().makeTranslation(x, y, 0);
            // Масштабируем
            let scale = Math.random() * (1 - 0.7) + 0.7;
            let ms = new THREE.Matrix4().makeScale(scale, scale, scale);
            // Крутим
            let mr = new THREE.Matrix4().makeRotationZ(Math.random() * (Math.PI / 2));
            // Формируем матрицу трансформации
            let m = new THREE.Matrix4().multiplyMatrices(mt, ms);
            m.multiply(mr);

            // Мержим деталь с геометрией
            geo.merge(geoDec, m);
        }

        let offset = this.terrain.verticesTopOffset;

        // Перебираем все тайлы чанка (от левого нижнего угла направо и вверх)
        for (let y = 0; y < this.terrain.chunkSize; y++) {

            // Вычисляем смещения индексов вершин
            let yb = y * (this.terrain.chunkSize + 1);
            let yt = yb + this.terrain.chunkSize + 1;

            for (let x = 0; x < this.terrain.chunkSize; x++) {

                // Индекс текущего тайла
                let index = x + y * this.terrain.chunkSize;

                // Вычисляем индексы 4 нижних точек тайла
                //
                //   0  *---*  1
                //      |   |
                //   2  *---*  3

                let i0 = yt + x;
                let i1 = yt + x + 1;
                let i2 = yb + x;
                let i3 = yb + x + 1;

                // Поверхность
                if (this.tiles[index].ground != 0) {
                    quad(geoLandscape, this.terrain.atlas, this.terrain.grounds[this.tiles[index].ground].texture, i0, i1, i2, i3);
                }

                // Деталь
                if (this.tiles[index].detail != 0) {
                    detail(geoDetails, this.terrain.atlas, this.terrain.details[this.tiles[index].detail].texture, x, y);
                };


                // Блок
                if (this.tiles[index].block != 0) {

                    // Южная сторона
                    quad(geoLandscape, this.terrain.atlas, this.terrain.blocks[this.tiles[index].block].textureWall, i2 + offset, i3 + offset, i2, i3);
                    // Восточная сторона
                    quad(geoLandscape, this.terrain.atlas, this.terrain.blocks[this.tiles[index].block].textureWall, i3 + offset, i1 + offset, i3, i1);
                    // Северная сторона
                    quad(geoLandscape, this.terrain.atlas, this.terrain.blocks[this.tiles[index].block].textureWall, i1 + offset, i0 + offset, i1, i0);
                    // Западная сторона
                    quad(geoLandscape, this.terrain.atlas, this.terrain.blocks[this.tiles[index].block].textureWall, i0 + offset, i2 + offset, i0, i2);
                    // Верхушка
                    quad(geoLandscape, this.terrain.atlas, this.terrain.blocks[this.tiles[index].block].textureTop, i0 + offset, i1 + offset, i2 + offset, i3 + offset);
                    
                }
            }
        }

        this.meshLandscape = new THREE.Mesh(geoLandscape, this.terrain.atlas.opaqueMaterial);
        this.meshLandscape.position.set(this.meshX, this.meshY, 0);

        this.meshDetails = new THREE.Mesh(geoDetails, this.terrain.atlas.transMaterial);
        this.meshDetails.position.set(this.meshX + 0.5, this.meshY + 0.5, 0);

    }

};

class Terrain {

    constructor(msgTerrain, atlas) {

        this.width = msgTerrain.width;
        this.height = msgTerrain.height;
        this.chunkSize = msgTerrain.chunkSize;
        this.grounds = msgTerrain.grounds;
        this.blocks = msgTerrain.blocks;
        this.details = msgTerrain.details;

        this.atlas = atlas;

        this.chunkedWidth = Math.ceil(this.width / this.chunkSize);
        this.chunkedHeight = Math.ceil(this.height / this.chunkSize);

        this.chunks = [];

        // Заполняем общий вершинный буфер для чанков
        this.verticesBuffer = [];
        for (let z = 0; z < 2; z++) {
            for (let y = 0; y < this.chunkSize + 1; y++) {
                for (let x = 0; x < this.chunkSize + 1; x++) {
                    this.verticesBuffer.push(new THREE.Vector3(x, y, z));
                }
            }
        }
        // Смещение индексов верхних вершин в буфере
        this.verticesTopOffset = (this.chunkSize + 1) * (this.chunkSize + 1);

    }

    setChunk(msgChunk) {

        this.chunks[msgChunk.index] = new Chunk(this, msgChunk);

    }

};


export { Terrain, Chunk };