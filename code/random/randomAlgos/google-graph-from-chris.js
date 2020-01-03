var $_$c = $_$wp(1);
const cave = ($_$w(1, 0, $_$c), [
    [
        [
            {
                blocked: false,
                treasure: false
            },
            {
                blocked: false,
                treasure: true
            }
        ],
        [
            {
                blocked: false,
                treasure: true
            },
            {
                blocked: false,
                treasure: false
            }
        ]
    ],
    [
        [
            {
                blocked: false,
                treasure: false
            },
            {
                blocked: false,
                treasure: false
            }
        ],
        [
            {
                blocked: false,
                treasure: true
            },
            {
                blocked: false,
                treasure: false
            }
        ]
    ]
]);
function Submersible(caveInput, startingX, startingY, startingZ) {
    var $_$c = $_$wf(1);
    $_$w(1, 1, $_$c), this.cave = caveInput;
    $_$w(1, 2, $_$c), this.x = startingX;
    $_$w(1, 3, $_$c), this.y = startingY;
    $_$w(1, 4, $_$c), this.z = startingZ;
    $_$w(1, 5, $_$c), this.treasuresFound = [];
    $_$w(1, 6, $_$c), this.currentPath = [];
}
$_$w(1, 7, $_$c), Submersible.prototype.scan = function scan() {
    var $_$c = $_$wf(1);
    const validDirections = ($_$w(1, 8, $_$c), []);
    const dx = ($_$w(1, 9, $_$c), [
        0,
        0,
        0,
        0,
        1,
        -1
    ]);
    const dy = ($_$w(1, 10, $_$c), [
        0,
        0,
        1,
        -1,
        0,
        0
    ]);
    const dz = ($_$w(1, 11, $_$c), [
        1,
        -1,
        0,
        0,
        0,
        0
    ]);
    for (let i = 0; $_$w(1, 12, $_$c), i < 6; i++) {
        const newX = ($_$w(1, 13, $_$c), this.x + dx[i]);
        const newY = ($_$w(1, 14, $_$c), this.y + dy[i]);
        const newZ = ($_$w(1, 15, $_$c), this.z + dz[i]);
        if ($_$w(1, 16, $_$c), ($_$w(1, 17, $_$c), ($_$w(1, 19, $_$c), ($_$w(1, 21, $_$c), ($_$w(1, 23, $_$c), ($_$w(1, 25, $_$c), newX >= 0) && ($_$w(1, 26, $_$c), newY >= 0)) && ($_$w(1, 24, $_$c), newZ >= 0)) && ($_$w(1, 22, $_$c), newX < this.cave.length)) && ($_$w(1, 20, $_$c), newY < this.cave[0].length)) && ($_$w(1, 18, $_$c), newZ < this.cave[0][0].length)) {
            if ($_$w(1, 27, $_$c), !cave[newX][newY][newZ].blocked) {
                $_$w(1, 28, $_$c), validDirections.push([
                    newX,
                    newY,
                    newZ
                ]);
            }
        }
    }
    return $_$w(1, 29, $_$c), validDirections;
};
$_$w(1, 30, $_$c), Submersible.prototype.investigate = function investigate() {
    var $_$c = $_$wf(1);
    return $_$w(1, 31, $_$c), this.cave[this.x][this.y][this.z].treasure;
};
$_$w(1, 32, $_$c), Submersible.prototype.move = function move(nextX, nextY, nextZ) {
    var $_$c = $_$wf(1);
    if ($_$w(1, 33, $_$c), Math.abs(nextX - this.x + nextY - this.y + nextZ - this.z) !== 1) {
        $_$w(1, 34, $_$c), $_$tracer.log('invalid movement', '', 1, 34);
        return $_$w(1, 35, $_$c), false;
    }
    $_$w(1, 36, $_$c), this.x = nextX;
    $_$w(1, 37, $_$c), this.y = nextY;
    $_$w(1, 38, $_$c), this.z = nextZ;
    return $_$w(1, 39, $_$c), true;
};
const sub = ($_$w(1, 40, $_$c), new Submersible(cave, 0, 0, 0));
function start() {
    var $_$c = $_$wf(1);
    if ($_$w(1, 41, $_$c), sub.currentPath.length === 0) {
        $_$w(1, 42, $_$c), sub.currentPath.push([
            this.startingX,
            this.startingY,
            this.startingZ
        ]);
    }
    if ($_$w(1, 43, $_$c), sub.scan().length === 0) {
        $_$w(1, 44, $_$c), sub.currentPath[sub.x][sub.y][sub.z].blocked = true;
        return $_$w(1, 45, $_$c), null;
    }
    $_$w(1, 46, $_$c), sub.scan().forEach(newCoords => {
        var $_$c = $_$wf(1);
        if ($_$w(1, 47, $_$c), sub.currentPath.filter(coords => {
                var $_$c = $_$wf(1);
                return $_$w(1, 48, $_$c), ($_$w(1, 49, $_$c), ($_$w(1, 51, $_$c), coords[0] === newCoords[0]) && ($_$w(1, 52, $_$c), coords[1] === newCoords[1])) && ($_$w(1, 50, $_$c), coords[2] === newCoords[2]);
            }).length > 0) {
        } else {
            $_$w(1, 53, $_$c), sub.currentPath.push(newCoords);
            $_$w(1, 54, $_$c), sub.move(...newCoords);
            if ($_$w(1, 55, $_$c), ($_$w(1, 56, $_$c), sub.investigate()) && ($_$w(1, 57, $_$c), sub.treasuresFound.filter(coords => {
                    var $_$c = $_$wf(1);
                    return $_$w(1, 58, $_$c), ($_$w(1, 59, $_$c), ($_$w(1, 61, $_$c), coords[0] === sub.x) && ($_$w(1, 62, $_$c), coords[1] === sub.y)) && ($_$w(1, 60, $_$c), coords[2] === sub.z);
                }).length === 0)) {
                $_$w(1, 63, $_$c), sub.treasuresFound.push([
                    sub.x,
                    sub.y,
                    sub.z
                ]);
            }
            $_$w(1, 64, $_$c), start();
        }
    });
    $_$w(1, 65, $_$c), sub.move(...sub.currentPath.pop());
    return $_$w(1, 66, $_$c), sub.treasuresFound;
}
const treasuresFound = ($_$w(1, 67, $_$c), start());
$_$w(1, 68, $_$c), $_$tracer.log(treasuresFound, 'treasuresFound', 1, 68);
$_$wpe(1);