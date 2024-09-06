package main

// Libspatialite - Adds spatial SQL capabilities to SQLite
// Homepage: https://www.gaia-gis.it/fossil/libspatialite/index

import (
	"fmt"
	
	"os/exec"
)

func installLibspatialite() {
	// Método 1: Descargar y extraer .tar.gz
	libspatialite_tar_url := "https://www.gaia-gis.it/gaia-sins/libspatialite-sources/libspatialite-5.1.0.tar.gz"
	libspatialite_cmd_tar := exec.Command("curl", "-L", libspatialite_tar_url, "-o", "package.tar.gz")
	err := libspatialite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspatialite_zip_url := "https://www.gaia-gis.it/gaia-sins/libspatialite-sources/libspatialite-5.1.0.zip"
	libspatialite_cmd_zip := exec.Command("curl", "-L", libspatialite_zip_url, "-o", "package.zip")
	err = libspatialite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspatialite_bin_url := "https://www.gaia-gis.it/gaia-sins/libspatialite-sources/libspatialite-5.1.0.bin"
	libspatialite_cmd_bin := exec.Command("curl", "-L", libspatialite_bin_url, "-o", "binary.bin")
	err = libspatialite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspatialite_src_url := "https://www.gaia-gis.it/gaia-sins/libspatialite-sources/libspatialite-5.1.0.src.tar.gz"
	libspatialite_cmd_src := exec.Command("curl", "-L", libspatialite_src_url, "-o", "source.tar.gz")
	err = libspatialite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspatialite_cmd_direct := exec.Command("./binary")
	err = libspatialite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freexl")
	exec.Command("latte", "install", "freexl").Run()
	fmt.Println("Instalando dependencia: geos")
	exec.Command("latte", "install", "geos").Run()
	fmt.Println("Instalando dependencia: librttopo")
	exec.Command("latte", "install", "librttopo").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
