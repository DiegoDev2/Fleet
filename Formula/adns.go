package main

// Adns - C/C++ resolver library and DNS resolver utilities
// Homepage: https://www.chiark.greenend.org.uk/~ian/adns/

import (
	"fmt"
	
	"os/exec"
)

func installAdns() {
	// Método 1: Descargar y extraer .tar.gz
	adns_tar_url := "https://www.chiark.greenend.org.uk/~ian/adns/ftp/adns-1.6.1.tar.gz"
	adns_cmd_tar := exec.Command("curl", "-L", adns_tar_url, "-o", "package.tar.gz")
	err := adns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adns_zip_url := "https://www.chiark.greenend.org.uk/~ian/adns/ftp/adns-1.6.1.zip"
	adns_cmd_zip := exec.Command("curl", "-L", adns_zip_url, "-o", "package.zip")
	err = adns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adns_bin_url := "https://www.chiark.greenend.org.uk/~ian/adns/ftp/adns-1.6.1.bin"
	adns_cmd_bin := exec.Command("curl", "-L", adns_bin_url, "-o", "binary.bin")
	err = adns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adns_src_url := "https://www.chiark.greenend.org.uk/~ian/adns/ftp/adns-1.6.1.src.tar.gz"
	adns_cmd_src := exec.Command("curl", "-L", adns_src_url, "-o", "source.tar.gz")
	err = adns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adns_cmd_direct := exec.Command("./binary")
	err = adns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
