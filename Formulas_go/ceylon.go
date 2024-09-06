package main

// Ceylon - Programming language for writing large programs in teams
// Homepage: https://projects.eclipse.org/projects/technology.ceylon

import (
	"fmt"
	
	"os/exec"
)

func installCeylon() {
	// Método 1: Descargar y extraer .tar.gz
	ceylon_tar_url := "https://web.archive.org/web/20200623041941/https://downloads.ceylon-lang.org/cli/ceylon-1.3.3.zip"
	ceylon_cmd_tar := exec.Command("curl", "-L", ceylon_tar_url, "-o", "package.tar.gz")
	err := ceylon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ceylon_zip_url := "https://web.archive.org/web/20200623041941/https://downloads.ceylon-lang.org/cli/ceylon-1.3.3.zip"
	ceylon_cmd_zip := exec.Command("curl", "-L", ceylon_zip_url, "-o", "package.zip")
	err = ceylon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ceylon_bin_url := "https://web.archive.org/web/20200623041941/https://downloads.ceylon-lang.org/cli/ceylon-1.3.3.zip"
	ceylon_cmd_bin := exec.Command("curl", "-L", ceylon_bin_url, "-o", "binary.bin")
	err = ceylon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ceylon_src_url := "https://web.archive.org/web/20200623041941/https://downloads.ceylon-lang.org/cli/ceylon-1.3.3.zip"
	ceylon_cmd_src := exec.Command("curl", "-L", ceylon_src_url, "-o", "source.tar.gz")
	err = ceylon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ceylon_cmd_direct := exec.Command("./binary")
	err = ceylon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
exec.Command("latte", "install", "openjdk@8")
}
