package main

// Launchdns - Mini DNS server designed solely to route queries to localhost
// Homepage: https://github.com/josh/launchdns

import (
	"fmt"
	
	"os/exec"
)

func installLaunchdns() {
	// Método 1: Descargar y extraer .tar.gz
	launchdns_tar_url := "https://github.com/josh/launchdns/archive/refs/tags/v1.0.4.tar.gz"
	launchdns_cmd_tar := exec.Command("curl", "-L", launchdns_tar_url, "-o", "package.tar.gz")
	err := launchdns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	launchdns_zip_url := "https://github.com/josh/launchdns/archive/refs/tags/v1.0.4.zip"
	launchdns_cmd_zip := exec.Command("curl", "-L", launchdns_zip_url, "-o", "package.zip")
	err = launchdns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	launchdns_bin_url := "https://github.com/josh/launchdns/archive/refs/tags/v1.0.4.bin"
	launchdns_cmd_bin := exec.Command("curl", "-L", launchdns_bin_url, "-o", "binary.bin")
	err = launchdns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	launchdns_src_url := "https://github.com/josh/launchdns/archive/refs/tags/v1.0.4.src.tar.gz"
	launchdns_cmd_src := exec.Command("curl", "-L", launchdns_src_url, "-o", "source.tar.gz")
	err = launchdns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	launchdns_cmd_direct := exec.Command("./binary")
	err = launchdns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
