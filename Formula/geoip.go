package main

// Geoip - This library is for the GeoIP Legacy format (dat)
// Homepage: https://github.com/maxmind/geoip-api-c

import (
	"fmt"
	
	"os/exec"
)

func installGeoip() {
	// Método 1: Descargar y extraer .tar.gz
	geoip_tar_url := "https://github.com/maxmind/geoip-api-c/releases/download/v1.6.12/GeoIP-1.6.12.tar.gz"
	geoip_cmd_tar := exec.Command("curl", "-L", geoip_tar_url, "-o", "package.tar.gz")
	err := geoip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geoip_zip_url := "https://github.com/maxmind/geoip-api-c/releases/download/v1.6.12/GeoIP-1.6.12.zip"
	geoip_cmd_zip := exec.Command("curl", "-L", geoip_zip_url, "-o", "package.zip")
	err = geoip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geoip_bin_url := "https://github.com/maxmind/geoip-api-c/releases/download/v1.6.12/GeoIP-1.6.12.bin"
	geoip_cmd_bin := exec.Command("curl", "-L", geoip_bin_url, "-o", "binary.bin")
	err = geoip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geoip_src_url := "https://github.com/maxmind/geoip-api-c/releases/download/v1.6.12/GeoIP-1.6.12.src.tar.gz"
	geoip_cmd_src := exec.Command("curl", "-L", geoip_src_url, "-o", "source.tar.gz")
	err = geoip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geoip_cmd_direct := exec.Command("./binary")
	err = geoip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
