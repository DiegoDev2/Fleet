package main

// Geoserver - Java server to share and edit geospatial data
// Homepage: https://geoserver.org/

import (
	"fmt"
	
	"os/exec"
)

func installGeoserver() {
	// Método 1: Descargar y extraer .tar.gz
	geoserver_tar_url := "https://downloads.sourceforge.net/project/geoserver/GeoServer/2.25.3/geoserver-2.25.3-bin.zip"
	geoserver_cmd_tar := exec.Command("curl", "-L", geoserver_tar_url, "-o", "package.tar.gz")
	err := geoserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geoserver_zip_url := "https://downloads.sourceforge.net/project/geoserver/GeoServer/2.25.3/geoserver-2.25.3-bin.zip"
	geoserver_cmd_zip := exec.Command("curl", "-L", geoserver_zip_url, "-o", "package.zip")
	err = geoserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geoserver_bin_url := "https://downloads.sourceforge.net/project/geoserver/GeoServer/2.25.3/geoserver-2.25.3-bin.zip"
	geoserver_cmd_bin := exec.Command("curl", "-L", geoserver_bin_url, "-o", "binary.bin")
	err = geoserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geoserver_src_url := "https://downloads.sourceforge.net/project/geoserver/GeoServer/2.25.3/geoserver-2.25.3-bin.zip"
	geoserver_cmd_src := exec.Command("curl", "-L", geoserver_src_url, "-o", "source.tar.gz")
	err = geoserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geoserver_cmd_direct := exec.Command("./binary")
	err = geoserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
