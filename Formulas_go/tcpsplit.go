package main

// Tcpsplit - Break a packet trace into some number of sub-traces
// Homepage: https://web.archive.org/web/20230609122227/https://www.icir.org/mallman/software/tcpsplit/

import (
	"fmt"
	
	"os/exec"
)

func installTcpsplit() {
	// Método 1: Descargar y extraer .tar.gz
	tcpsplit_tar_url := "https://web.archive.org/web/20230609122227/https://www.icir.org/mallman/software/tcpsplit/tcpsplit-0.2.tar.gz"
	tcpsplit_cmd_tar := exec.Command("curl", "-L", tcpsplit_tar_url, "-o", "package.tar.gz")
	err := tcpsplit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcpsplit_zip_url := "https://web.archive.org/web/20230609122227/https://www.icir.org/mallman/software/tcpsplit/tcpsplit-0.2.zip"
	tcpsplit_cmd_zip := exec.Command("curl", "-L", tcpsplit_zip_url, "-o", "package.zip")
	err = tcpsplit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcpsplit_bin_url := "https://web.archive.org/web/20230609122227/https://www.icir.org/mallman/software/tcpsplit/tcpsplit-0.2.bin"
	tcpsplit_cmd_bin := exec.Command("curl", "-L", tcpsplit_bin_url, "-o", "binary.bin")
	err = tcpsplit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcpsplit_src_url := "https://web.archive.org/web/20230609122227/https://www.icir.org/mallman/software/tcpsplit/tcpsplit-0.2.src.tar.gz"
	tcpsplit_cmd_src := exec.Command("curl", "-L", tcpsplit_src_url, "-o", "source.tar.gz")
	err = tcpsplit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcpsplit_cmd_direct := exec.Command("./binary")
	err = tcpsplit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
