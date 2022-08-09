/* Initial Map */
var map = L.map("map").setView([-7.5418696, 110.4462433], 10);

/* Tile Basemap */
var basemap = L.tileLayer(
	"https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
	{
		attribution:
			'<a href="https://unsorry.net" target="_blank">unsorry@2022</a>',
	}
);
basemap.addTo(map);

map.attributionControl.setPrefix(false);

/* Marker */
var marker = L.marker([-7.5418696, 110.4462433], { draggable: true });
marker.addTo(map);
marker.bindPopup("Geser Marker");
marker.openPopup();

marker.on("dragend", function (e) {
	var latlng = e.target.getLatLng();
	map.flyTo(new L.LatLng(latlng.lat.toFixed(9), latlng.lng.toFixed(9)));
	marker.bindPopup(
		"Koordinat:<br>" + latlng.lat.toFixed(7) + "," + latlng.lng.toFixed(7)
	);
	marker.openPopup();
});

/* GeoJSON Point */
var titikkabkota = L.geoJson(null, {
	pointToLayer: function (feature, latlng) {
		return L.marker(latlng, {
			icon: L.icon({
				iconUrl: "/static/image/city-marker.png", //icon simbol
				iconSize: [36, 36], //ukuran icon simbol
				iconAnchor: [18, 36], //penempatan icon simbol
				popupAnchor: [0, -36], //penempatan popup terhadap icon simbol
			}),
		});
	},
	onEachFeature: function (feature, layer) {
		if (feature.properties) {
			var content =
				"Koordinat: " +
				feature.geometry.coordinates[1] +
				", " +
				feature.geometry.coordinates[0];
			layer.on({
				click: function (e) {
					//Fungsi ketika icon simbol di-klik
					e.bindPopup(content);
				},
				mouseover: function (e) {
					e.bindTooltip(content);
				},
			});
		}
	},
});
/* memanggil data geojson point */
$.getJSON("/api", function (data) {
	titikkabkota.addData(data);
	map.addLayer(titikkabkota); //titikkabkota ditampilkan ketika halaman dipanggil
});
