/* Initial Map */
var map = L.map('map').setView([-7.9,110.45],10);
    
/* Tile Basemap */
var basemap = L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
	attribution: '<a href="https://unsorry.net" target="_blank">unsorry@2022</a>'
});
basemap.addTo(map);

map.attributionControl.setPrefix(false)

/* Marker */
var marker = L.marker([-7.8,110.37], {draggable: true});
marker.addTo(map);
marker.bindPopup('Geser Marker');
marker.openPopup();

marker.on('dragend', function (e) {
	var latlng = e.target.getLatLng();
	map.flyTo(new L.LatLng(latlng.lat.toFixed(9), latlng.lng.toFixed(9)));
	marker.bindPopup('Koordinat:<br>' + latlng.lat.toFixed(7) + ',' + latlng.lng.toFixed(7));
	marker.openPopup();
});