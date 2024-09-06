package main

// Syck - Extension for reading and writing YAML
// Homepage: https://github.com/indeyets/syck

import (
	"fmt"
	
	"os/exec"
)

func installSyck() {
	// Método 1: Descargar y extraer .tar.gz
	syck_tar_url := "https://github.s3.amazonaws.com/downloads/indeyets/syck/syck-0.70.tar.gz"
	syck_cmd_tar := exec.Command("curl", "-L", syck_tar_url, "-o", "package.tar.gz")
	err := syck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syck_zip_url := "https://github.s3.amazonaws.com/downloads/indeyets/syck/syck-0.70.zip"
	syck_cmd_zip := exec.Command("curl", "-L", syck_zip_url, "-o", "package.zip")
	err = syck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syck_bin_url := "https://github.s3.amazonaws.com/downloads/indeyets/syck/syck-0.70.bin"
	syck_cmd_bin := exec.Command("curl", "-L", syck_bin_url, "-o", "binary.bin")
	err = syck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syck_src_url := "https://github.s3.amazonaws.com/downloads/indeyets/syck/syck-0.70.src.tar.gz"
	syck_cmd_src := exec.Command("curl", "-L", syck_src_url, "-o", "source.tar.gz")
	err = syck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syck_cmd_direct := exec.Command("./binary")
	err = syck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
