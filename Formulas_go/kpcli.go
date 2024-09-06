package main

// Kpcli - Command-line interface to KeePass database files
// Homepage: https://kpcli.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installKpcli() {
	// Método 1: Descargar y extraer .tar.gz
	kpcli_tar_url := "https://downloads.sourceforge.net/project/kpcli/kpcli-4.1.pl"
	kpcli_cmd_tar := exec.Command("curl", "-L", kpcli_tar_url, "-o", "package.tar.gz")
	err := kpcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kpcli_zip_url := "https://downloads.sourceforge.net/project/kpcli/kpcli-4.1.pl"
	kpcli_cmd_zip := exec.Command("curl", "-L", kpcli_zip_url, "-o", "package.zip")
	err = kpcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kpcli_bin_url := "https://downloads.sourceforge.net/project/kpcli/kpcli-4.1.pl"
	kpcli_cmd_bin := exec.Command("curl", "-L", kpcli_bin_url, "-o", "binary.bin")
	err = kpcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kpcli_src_url := "https://downloads.sourceforge.net/project/kpcli/kpcli-4.1.pl"
	kpcli_cmd_src := exec.Command("curl", "-L", kpcli_src_url, "-o", "source.tar.gz")
	err = kpcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kpcli_cmd_direct := exec.Command("./binary")
	err = kpcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
