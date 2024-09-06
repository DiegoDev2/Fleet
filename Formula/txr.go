package main

// Txr - Lisp-like programming language for convenient data munging
// Homepage: https://www.nongnu.org/txr/

import (
	"fmt"
	
	"os/exec"
)

func installTxr() {
	// Método 1: Descargar y extraer .tar.gz
	txr_tar_url := "https://www.kylheku.com/cgit/txr/snapshot/txr-296.tar.bz2"
	txr_cmd_tar := exec.Command("curl", "-L", txr_tar_url, "-o", "package.tar.gz")
	err := txr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	txr_zip_url := "https://www.kylheku.com/cgit/txr/snapshot/txr-296.tar.bz2"
	txr_cmd_zip := exec.Command("curl", "-L", txr_zip_url, "-o", "package.zip")
	err = txr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	txr_bin_url := "https://www.kylheku.com/cgit/txr/snapshot/txr-296.tar.bz2"
	txr_cmd_bin := exec.Command("curl", "-L", txr_bin_url, "-o", "binary.bin")
	err = txr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	txr_src_url := "https://www.kylheku.com/cgit/txr/snapshot/txr-296.tar.bz2"
	txr_cmd_src := exec.Command("curl", "-L", txr_src_url, "-o", "source.tar.gz")
	err = txr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	txr_cmd_direct := exec.Command("./binary")
	err = txr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
