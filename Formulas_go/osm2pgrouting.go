package main

// Osm2pgrouting - Import OSM data into pgRouting database
// Homepage: https://pgrouting.org/docs/tools/osm2pgrouting.html

import (
	"fmt"
	
	"os/exec"
)

func installOsm2pgrouting() {
	// Método 1: Descargar y extraer .tar.gz
	osm2pgrouting_tar_url := "https://github.com/pgRouting/osm2pgrouting/archive/refs/tags/v2.3.8.tar.gz"
	osm2pgrouting_cmd_tar := exec.Command("curl", "-L", osm2pgrouting_tar_url, "-o", "package.tar.gz")
	err := osm2pgrouting_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osm2pgrouting_zip_url := "https://github.com/pgRouting/osm2pgrouting/archive/refs/tags/v2.3.8.zip"
	osm2pgrouting_cmd_zip := exec.Command("curl", "-L", osm2pgrouting_zip_url, "-o", "package.zip")
	err = osm2pgrouting_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osm2pgrouting_bin_url := "https://github.com/pgRouting/osm2pgrouting/archive/refs/tags/v2.3.8.bin"
	osm2pgrouting_cmd_bin := exec.Command("curl", "-L", osm2pgrouting_bin_url, "-o", "binary.bin")
	err = osm2pgrouting_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osm2pgrouting_src_url := "https://github.com/pgRouting/osm2pgrouting/archive/refs/tags/v2.3.8.src.tar.gz"
	osm2pgrouting_cmd_src := exec.Command("curl", "-L", osm2pgrouting_src_url, "-o", "source.tar.gz")
	err = osm2pgrouting_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osm2pgrouting_cmd_direct := exec.Command("./binary")
	err = osm2pgrouting_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: libpqxx")
exec.Command("latte", "install", "libpqxx")
	fmt.Println("Instalando dependencia: pgrouting")
exec.Command("latte", "install", "pgrouting")
	fmt.Println("Instalando dependencia: postgis")
exec.Command("latte", "install", "postgis")
}
