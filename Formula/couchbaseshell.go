package main

// CouchbaseShell - Modern and fun shell for Couchbase Server and Capella
// Homepage: https://couchbase.sh

import (
	"fmt"
	
	"os/exec"
)

func installCouchbaseShell() {
	// Método 1: Descargar y extraer .tar.gz
	couchbaseshell_tar_url := "https://github.com/couchbaselabs/couchbase-shell/archive/refs/tags/v0.75.2.tar.gz"
	couchbaseshell_cmd_tar := exec.Command("curl", "-L", couchbaseshell_tar_url, "-o", "package.tar.gz")
	err := couchbaseshell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	couchbaseshell_zip_url := "https://github.com/couchbaselabs/couchbase-shell/archive/refs/tags/v0.75.2.zip"
	couchbaseshell_cmd_zip := exec.Command("curl", "-L", couchbaseshell_zip_url, "-o", "package.zip")
	err = couchbaseshell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	couchbaseshell_bin_url := "https://github.com/couchbaselabs/couchbase-shell/archive/refs/tags/v0.75.2.bin"
	couchbaseshell_cmd_bin := exec.Command("curl", "-L", couchbaseshell_bin_url, "-o", "binary.bin")
	err = couchbaseshell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	couchbaseshell_src_url := "https://github.com/couchbaselabs/couchbase-shell/archive/refs/tags/v0.75.2.src.tar.gz"
	couchbaseshell_cmd_src := exec.Command("curl", "-L", couchbaseshell_src_url, "-o", "source.tar.gz")
	err = couchbaseshell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	couchbaseshell_cmd_direct := exec.Command("./binary")
	err = couchbaseshell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm@15")
	exec.Command("latte", "install", "llvm@15").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
