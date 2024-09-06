package main

// SpatialiteTools - CLI tools supporting SpatiaLite
// Homepage: https://www.gaia-gis.it/fossil/spatialite-tools/index

import (
	"fmt"
	
	"os/exec"
)

func installSpatialiteTools() {
	// Método 1: Descargar y extraer .tar.gz
	spatialitetools_tar_url := "https://www.gaia-gis.it/gaia-sins/spatialite-tools-sources/spatialite-tools-5.1.0.tar.gz"
	spatialitetools_cmd_tar := exec.Command("curl", "-L", spatialitetools_tar_url, "-o", "package.tar.gz")
	err := spatialitetools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spatialitetools_zip_url := "https://www.gaia-gis.it/gaia-sins/spatialite-tools-sources/spatialite-tools-5.1.0.zip"
	spatialitetools_cmd_zip := exec.Command("curl", "-L", spatialitetools_zip_url, "-o", "package.zip")
	err = spatialitetools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spatialitetools_bin_url := "https://www.gaia-gis.it/gaia-sins/spatialite-tools-sources/spatialite-tools-5.1.0.bin"
	spatialitetools_cmd_bin := exec.Command("curl", "-L", spatialitetools_bin_url, "-o", "binary.bin")
	err = spatialitetools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spatialitetools_src_url := "https://www.gaia-gis.it/gaia-sins/spatialite-tools-sources/spatialite-tools-5.1.0.src.tar.gz"
	spatialitetools_cmd_src := exec.Command("curl", "-L", spatialitetools_src_url, "-o", "source.tar.gz")
	err = spatialitetools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spatialitetools_cmd_direct := exec.Command("./binary")
	err = spatialitetools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libspatialite")
exec.Command("latte", "install", "libspatialite")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
	fmt.Println("Instalando dependencia: readosm")
exec.Command("latte", "install", "readosm")
}
