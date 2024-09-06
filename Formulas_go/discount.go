package main

// Discount - C implementation of Markdown
// Homepage: https://www.pell.portland.or.us/~orc/Code/discount/

import (
	"fmt"
	
	"os/exec"
)

func installDiscount() {
	// Método 1: Descargar y extraer .tar.gz
	discount_tar_url := "https://www.pell.portland.or.us/~orc/Code/discount/discount-2.2.7d.tar.bz2"
	discount_cmd_tar := exec.Command("curl", "-L", discount_tar_url, "-o", "package.tar.gz")
	err := discount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	discount_zip_url := "https://www.pell.portland.or.us/~orc/Code/discount/discount-2.2.7d.tar.bz2"
	discount_cmd_zip := exec.Command("curl", "-L", discount_zip_url, "-o", "package.zip")
	err = discount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	discount_bin_url := "https://www.pell.portland.or.us/~orc/Code/discount/discount-2.2.7d.tar.bz2"
	discount_cmd_bin := exec.Command("curl", "-L", discount_bin_url, "-o", "binary.bin")
	err = discount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	discount_src_url := "https://www.pell.portland.or.us/~orc/Code/discount/discount-2.2.7d.tar.bz2"
	discount_cmd_src := exec.Command("curl", "-L", discount_src_url, "-o", "source.tar.gz")
	err = discount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	discount_cmd_direct := exec.Command("./binary")
	err = discount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
