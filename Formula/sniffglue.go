package main

// Sniffglue - Secure multithreaded packet sniffer
// Homepage: https://github.com/kpcyrd/sniffglue

import (
	"fmt"
	
	"os/exec"
)

func installSniffglue() {
	// Método 1: Descargar y extraer .tar.gz
	sniffglue_tar_url := "https://github.com/kpcyrd/sniffglue/archive/refs/tags/v0.16.0.tar.gz"
	sniffglue_cmd_tar := exec.Command("curl", "-L", sniffglue_tar_url, "-o", "package.tar.gz")
	err := sniffglue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sniffglue_zip_url := "https://github.com/kpcyrd/sniffglue/archive/refs/tags/v0.16.0.zip"
	sniffglue_cmd_zip := exec.Command("curl", "-L", sniffglue_zip_url, "-o", "package.zip")
	err = sniffglue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sniffglue_bin_url := "https://github.com/kpcyrd/sniffglue/archive/refs/tags/v0.16.0.bin"
	sniffglue_cmd_bin := exec.Command("curl", "-L", sniffglue_bin_url, "-o", "binary.bin")
	err = sniffglue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sniffglue_src_url := "https://github.com/kpcyrd/sniffglue/archive/refs/tags/v0.16.0.src.tar.gz"
	sniffglue_cmd_src := exec.Command("curl", "-L", sniffglue_src_url, "-o", "source.tar.gz")
	err = sniffglue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sniffglue_cmd_direct := exec.Command("./binary")
	err = sniffglue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
}
