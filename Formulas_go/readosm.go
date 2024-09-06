package main

// Readosm - Extract valid data from an Open Street Map input file
// Homepage: https://www.gaia-gis.it/fossil/readosm/index

import (
	"fmt"
	
	"os/exec"
)

func installReadosm() {
	// Método 1: Descargar y extraer .tar.gz
	readosm_tar_url := "https://www.gaia-gis.it/gaia-sins/readosm-sources/readosm-1.1.0a.tar.gz"
	readosm_cmd_tar := exec.Command("curl", "-L", readosm_tar_url, "-o", "package.tar.gz")
	err := readosm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	readosm_zip_url := "https://www.gaia-gis.it/gaia-sins/readosm-sources/readosm-1.1.0a.zip"
	readosm_cmd_zip := exec.Command("curl", "-L", readosm_zip_url, "-o", "package.zip")
	err = readosm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	readosm_bin_url := "https://www.gaia-gis.it/gaia-sins/readosm-sources/readosm-1.1.0a.bin"
	readosm_cmd_bin := exec.Command("curl", "-L", readosm_bin_url, "-o", "binary.bin")
	err = readosm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	readosm_src_url := "https://www.gaia-gis.it/gaia-sins/readosm-sources/readosm-1.1.0a.src.tar.gz"
	readosm_cmd_src := exec.Command("curl", "-L", readosm_src_url, "-o", "source.tar.gz")
	err = readosm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	readosm_cmd_direct := exec.Command("./binary")
	err = readosm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
