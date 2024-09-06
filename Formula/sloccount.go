package main

// Sloccount - Count lines of code in many languages
// Homepage: https://dwheeler.com/sloccount/

import (
	"fmt"
	
	"os/exec"
)

func installSloccount() {
	// Método 1: Descargar y extraer .tar.gz
	sloccount_tar_url := "https://dwheeler.com/sloccount/sloccount-2.26.tar.gz"
	sloccount_cmd_tar := exec.Command("curl", "-L", sloccount_tar_url, "-o", "package.tar.gz")
	err := sloccount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sloccount_zip_url := "https://dwheeler.com/sloccount/sloccount-2.26.zip"
	sloccount_cmd_zip := exec.Command("curl", "-L", sloccount_zip_url, "-o", "package.zip")
	err = sloccount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sloccount_bin_url := "https://dwheeler.com/sloccount/sloccount-2.26.bin"
	sloccount_cmd_bin := exec.Command("curl", "-L", sloccount_bin_url, "-o", "binary.bin")
	err = sloccount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sloccount_src_url := "https://dwheeler.com/sloccount/sloccount-2.26.src.tar.gz"
	sloccount_cmd_src := exec.Command("curl", "-L", sloccount_src_url, "-o", "source.tar.gz")
	err = sloccount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sloccount_cmd_direct := exec.Command("./binary")
	err = sloccount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
