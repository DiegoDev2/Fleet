package main

// Osm2pgsql - OpenStreetMap data to PostgreSQL converter
// Homepage: https://osm2pgsql.org

import (
	"fmt"
	
	"os/exec"
)

func installOsm2pgsql() {
	// Método 1: Descargar y extraer .tar.gz
	osm2pgsql_tar_url := "https://github.com/openstreetmap/osm2pgsql/archive/refs/tags/1.11.0.tar.gz"
	osm2pgsql_cmd_tar := exec.Command("curl", "-L", osm2pgsql_tar_url, "-o", "package.tar.gz")
	err := osm2pgsql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osm2pgsql_zip_url := "https://github.com/openstreetmap/osm2pgsql/archive/refs/tags/1.11.0.zip"
	osm2pgsql_cmd_zip := exec.Command("curl", "-L", osm2pgsql_zip_url, "-o", "package.zip")
	err = osm2pgsql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osm2pgsql_bin_url := "https://github.com/openstreetmap/osm2pgsql/archive/refs/tags/1.11.0.bin"
	osm2pgsql_cmd_bin := exec.Command("curl", "-L", osm2pgsql_bin_url, "-o", "binary.bin")
	err = osm2pgsql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osm2pgsql_src_url := "https://github.com/openstreetmap/osm2pgsql/archive/refs/tags/1.11.0.src.tar.gz"
	osm2pgsql_cmd_src := exec.Command("curl", "-L", osm2pgsql_src_url, "-o", "source.tar.gz")
	err = osm2pgsql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osm2pgsql_cmd_direct := exec.Command("./binary")
	err = osm2pgsql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
}
