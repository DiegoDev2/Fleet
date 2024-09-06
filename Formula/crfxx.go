package main

// Crfxx - Conditional random fields for segmenting/labeling sequential data
// Homepage: https://taku910.github.io/crfpp/

import (
	"fmt"
	
	"os/exec"
)

func installCrfxx() {
	// Método 1: Descargar y extraer .tar.gz
	crfxx_tar_url := "https://mirrors.sohu.com/gentoo/distfiles/f2/CRF%2B%2B-0.58.tar.gz"
	crfxx_cmd_tar := exec.Command("curl", "-L", crfxx_tar_url, "-o", "package.tar.gz")
	err := crfxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crfxx_zip_url := "https://mirrors.sohu.com/gentoo/distfiles/f2/CRF%2B%2B-0.58.zip"
	crfxx_cmd_zip := exec.Command("curl", "-L", crfxx_zip_url, "-o", "package.zip")
	err = crfxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crfxx_bin_url := "https://mirrors.sohu.com/gentoo/distfiles/f2/CRF%2B%2B-0.58.bin"
	crfxx_cmd_bin := exec.Command("curl", "-L", crfxx_bin_url, "-o", "binary.bin")
	err = crfxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crfxx_src_url := "https://mirrors.sohu.com/gentoo/distfiles/f2/CRF%2B%2B-0.58.src.tar.gz"
	crfxx_cmd_src := exec.Command("curl", "-L", crfxx_src_url, "-o", "source.tar.gz")
	err = crfxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crfxx_cmd_direct := exec.Command("./binary")
	err = crfxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
