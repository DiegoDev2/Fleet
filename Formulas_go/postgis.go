package main

// Postgis - Adds support for geographic objects to PostgreSQL
// Homepage: https://postgis.net/

import (
	"fmt"
	
	"os/exec"
)

func installPostgis() {
	// Método 1: Descargar y extraer .tar.gz
	postgis_tar_url := "https://download.osgeo.org/postgis/source/postgis-3.4.3.tar.gz"
	postgis_cmd_tar := exec.Command("curl", "-L", postgis_tar_url, "-o", "package.tar.gz")
	err := postgis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	postgis_zip_url := "https://download.osgeo.org/postgis/source/postgis-3.4.3.zip"
	postgis_cmd_zip := exec.Command("curl", "-L", postgis_zip_url, "-o", "package.zip")
	err = postgis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	postgis_bin_url := "https://download.osgeo.org/postgis/source/postgis-3.4.3.bin"
	postgis_cmd_bin := exec.Command("curl", "-L", postgis_bin_url, "-o", "binary.bin")
	err = postgis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	postgis_src_url := "https://download.osgeo.org/postgis/source/postgis-3.4.3.src.tar.gz"
	postgis_cmd_src := exec.Command("curl", "-L", postgis_src_url, "-o", "source.tar.gz")
	err = postgis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	postgis_cmd_direct := exec.Command("./binary")
	err = postgis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gpp")
exec.Command("latte", "install", "gpp")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gdal")
exec.Command("latte", "install", "gdal")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: postgresql@14")
exec.Command("latte", "install", "postgresql@14")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: protobuf-c")
exec.Command("latte", "install", "protobuf-c")
	fmt.Println("Instalando dependencia: sfcgal")
exec.Command("latte", "install", "sfcgal")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
