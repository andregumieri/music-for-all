var data = [
    {
        'artist': 'Chitãozinho & Xororó', 
        'albuns': [
            {'spotify_id': '2zK703NgF1doU9Rdsfwwrp', 'cover': 'https://4.bp.blogspot.com/-ab-spt9RdEg/Tpi3XGkQh7I/AAAAAAAABhs/AQzqN9kOX6o/s1600/chitaozinho+e+xoxoro+-+fio+de+cabelo+%2528frente.%2529.jpg'},
            {'spotify_id': '3j9uI8HhDZuljfucwyOUJ3', 'cover': 'https://i.scdn.co/image/ab67616d0000b273030ac523f895d929f5ba1bd8'},
            {'spotify_id': '1c3tyuPzbjc2fHmLtQ8tsK', 'cover': 'https://i.scdn.co/image/ab67616d0000b2731581e8be31c598de03bcb6d7'},
            {'spotify_id': '5BUJNkKh6wa1uxsMq7NAFR', 'cover': 'https://i.scdn.co/image/ab67616d0000b273aed75fe7a7ff4f214ca7aa34'},
            {'spotify_id': '5zNBpMWo4DajRLQ5MNolcO', 'cover': 'https://i.scdn.co/image/ab67616d0000b2735d44fc80e171a4a5d7de38df'},
        ]
    },

    {
        'artist': 'Zezé di Camargo & Luciano', 
        'albuns': [
            {'spotify_id': '5XsicqXM0n9m7gp9IXD8kO', 'cover': 'https://i.scdn.co/image/ab67616d00001e0246401faa35b9a976385d8375'},
            {'spotify_id': '6ORajOjreYSAI4tv4DJA4f', 'cover': 'https://i.scdn.co/image/ab67616d0000b273f40fc73b5db2db8d18b466a6'},
            {'spotify_id': '6Ls690GTjzbU4YNHnyCg1M', 'cover': 'https://i.scdn.co/image/ab67616d0000b273f5b0d511b5c1669f800c88fe'},
        ]
    },

    {
        'artist': 'Leandro & Leonardo', 
        'albuns': [
            {'spotify_id': '4wf7wWhbh5bPpPCG2YbvuV', 'cover': 'https://i.scdn.co/image/ab67616d0000b2736587dcf9d386227d272aecef'},
            {'spotify_id': '61MD8Pz2Wa5Ap1dOZc41fI', 'cover': 'https://i.scdn.co/image/ab67616d0000b273e0cb91cc709af0823d623961'},
            {'spotify_id': '7IlOlrjDC9wKYPWndzekqF', 'cover': 'https://i.scdn.co/image/ab67616d0000b273d3a9c1f66f2e97181314195a'},
            {'spotify_id': '0k0xdEdiB9kVIjnkXj1Qut', 'cover': 'https://i.scdn.co/image/ab67616d0000b2731195186fceab7af01ba0d51f'},
            {'spotify_id': '0adc2tR6B25LntLqQkGStj', 'cover': 'https://i.scdn.co/image/ab67616d0000b27338b58e2d33a2550e5e19f472'},
        ]
    },

    {
        'artist': 'Amigos', 
        'albuns': [
            {'spotify_id': '1k8op45qU1UrER9X0R8tv0', 'cover': 'https://www.modaoms.com.br/wp-content/uploads/2019/12/500x500-000000-80-0-0-1.jpg'},
        ]
    },
]


var list = document.getElementById('artistsList');
var albunsWrap = document.getElementById('albunsWrap');

function populate(x) {
    albunsWrap.innerHTML = '';
    var artist = data[x];

    for(var y = 0; y < artist.albuns.length; y++) {
        var album = artist.albuns[y];
        var divCover = document.createElement('div');
        var iframe = document.createElement('iframe');

        divCover.style.backgroundImage = 'url(' + album.cover + ')';
        divCover.className='cover';
        iframe.frameBorder = '0';
        iframe.allowTransparency = 'true';
        iframe.allow = 'encrypted-media';
        iframe.src = 'https://open.spotify.com/embed/album/' + album.spotify_id;

        divCover.append(iframe);
        albunsWrap.append(divCover);
    }
}

function change() {
    document.querySelector('.btn.btn--active').className = document.querySelector('.btn.btn--active').className.replace('btn--active', '').trim();
    this.className += ' btn--active';
    populate(this._index);
}


/*Add Artist Section Header*/
var artistColumnTitle = document.createElement('div');
artistColumnTitle._index = -1;
artistColumnTitle.innerHTML = "Artists";
artistColumnTitle.className = 'btn';
artistColumnTitle.style.backgroundColor = "#eeeeee";
list.append(artistColumnTitle)

for(var x = 0; x < data.length; x++) {
    var artist = data[x];
    var btnArtist = document.createElement('a');
    var coverWrap = document.createElement('div');
    btnArtist._index = x;
    btnArtist.innerHTML = artist.artist;
    btnArtist.className = 'btn';
    list.append(btnArtist);

    if(x == 0) {
        btnArtist.className += ' btn--active';
    }

    btnArtist.addEventListener('click', change, false);
    
    populate(0);
}