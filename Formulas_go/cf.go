package main

// Cf - Filter to replace numeric timestamps with a formatted date time
// Homepage: https://ee.lbl.gov/

import (
	"fmt"
	
	"os/exec"
)

func installCf() {
	// Método 1: Descargar y extraer .tar.gz
	cf_tar_url := "https://ee.lbl.gov/downloads/cf/cf-1.2.8.tar.gz"
	cf_cmd_tar := exec.Command("curl", "-L", cf_tar_url, "-o", "package.tar.gz")
	err := cf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cf_zip_url := "https://ee.lbl.gov/downloads/cf/cf-1.2.8.zip"
	cf_cmd_zip := exec.Command("curl", "-L", cf_zip_url, "-o", "package.zip")
	err = cf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cf_bin_url := "https://ee.lbl.gov/downloads/cf/cf-1.2.8.bin"
	cf_cmd_bin := exec.Command("curl", "-L", cf_bin_url, "-o", "binary.bin")
	err = cf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cf_src_url := "https://ee.lbl.gov/downloads/cf/cf-1.2.8.src.tar.gz"
	cf_cmd_src := exec.Command("curl", "-L", cf_src_url, "-o", "source.tar.gz")
	err = cf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cf_cmd_direct := exec.Command("./binary")
	err = cf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
