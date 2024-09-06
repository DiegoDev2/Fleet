package main

// ProxychainsNg - Hook preloader
// Homepage: https://github.com/rofl0r/proxychains-ng

import (
	"fmt"
	
	"os/exec"
)

func installProxychainsNg() {
	// Método 1: Descargar y extraer .tar.gz
	proxychainsng_tar_url := "https://github.com/rofl0r/proxychains-ng/archive/refs/tags/v4.17.tar.gz"
	proxychainsng_cmd_tar := exec.Command("curl", "-L", proxychainsng_tar_url, "-o", "package.tar.gz")
	err := proxychainsng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proxychainsng_zip_url := "https://github.com/rofl0r/proxychains-ng/archive/refs/tags/v4.17.zip"
	proxychainsng_cmd_zip := exec.Command("curl", "-L", proxychainsng_zip_url, "-o", "package.zip")
	err = proxychainsng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proxychainsng_bin_url := "https://github.com/rofl0r/proxychains-ng/archive/refs/tags/v4.17.bin"
	proxychainsng_cmd_bin := exec.Command("curl", "-L", proxychainsng_bin_url, "-o", "binary.bin")
	err = proxychainsng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proxychainsng_src_url := "https://github.com/rofl0r/proxychains-ng/archive/refs/tags/v4.17.src.tar.gz"
	proxychainsng_cmd_src := exec.Command("curl", "-L", proxychainsng_src_url, "-o", "source.tar.gz")
	err = proxychainsng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proxychainsng_cmd_direct := exec.Command("./binary")
	err = proxychainsng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
