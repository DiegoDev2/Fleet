package main

// Xa - 6502 cross assembler
// Homepage: https://www.floodgap.com/retrotech/xa/

import (
	"fmt"
	
	"os/exec"
)

func installXa() {
	// Método 1: Descargar y extraer .tar.gz
	xa_tar_url := "https://www.floodgap.com/retrotech/xa/dists/xa-2.4.1.tar.gz"
	xa_cmd_tar := exec.Command("curl", "-L", xa_tar_url, "-o", "package.tar.gz")
	err := xa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xa_zip_url := "https://www.floodgap.com/retrotech/xa/dists/xa-2.4.1.zip"
	xa_cmd_zip := exec.Command("curl", "-L", xa_zip_url, "-o", "package.zip")
	err = xa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xa_bin_url := "https://www.floodgap.com/retrotech/xa/dists/xa-2.4.1.bin"
	xa_cmd_bin := exec.Command("curl", "-L", xa_bin_url, "-o", "binary.bin")
	err = xa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xa_src_url := "https://www.floodgap.com/retrotech/xa/dists/xa-2.4.1.src.tar.gz"
	xa_cmd_src := exec.Command("curl", "-L", xa_src_url, "-o", "source.tar.gz")
	err = xa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xa_cmd_direct := exec.Command("./binary")
	err = xa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
