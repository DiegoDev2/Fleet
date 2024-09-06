package main

// Bashdb - Bash shell debugger
// Homepage: https://bashdb.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBashdb() {
	// Método 1: Descargar y extraer .tar.gz
	bashdb_tar_url := "https://downloads.sourceforge.net/project/bashdb/bashdb/5.0-1.1.2/bashdb-5.0-1.1.2.tar.bz2"
	bashdb_cmd_tar := exec.Command("curl", "-L", bashdb_tar_url, "-o", "package.tar.gz")
	err := bashdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashdb_zip_url := "https://downloads.sourceforge.net/project/bashdb/bashdb/5.0-1.1.2/bashdb-5.0-1.1.2.tar.bz2"
	bashdb_cmd_zip := exec.Command("curl", "-L", bashdb_zip_url, "-o", "package.zip")
	err = bashdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashdb_bin_url := "https://downloads.sourceforge.net/project/bashdb/bashdb/5.0-1.1.2/bashdb-5.0-1.1.2.tar.bz2"
	bashdb_cmd_bin := exec.Command("curl", "-L", bashdb_bin_url, "-o", "binary.bin")
	err = bashdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashdb_src_url := "https://downloads.sourceforge.net/project/bashdb/bashdb/5.0-1.1.2/bashdb-5.0-1.1.2.tar.bz2"
	bashdb_cmd_src := exec.Command("curl", "-L", bashdb_src_url, "-o", "source.tar.gz")
	err = bashdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashdb_cmd_direct := exec.Command("./binary")
	err = bashdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
}
