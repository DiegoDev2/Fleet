package main

// Sgrep - Search SGML, XML, and HTML
// Homepage: https://www.cs.helsinki.fi/u/jjaakkol/sgrep.html

import (
	"fmt"
	
	"os/exec"
)

func installSgrep() {
	// Método 1: Descargar y extraer .tar.gz
	sgrep_tar_url := "https://deb.debian.org/debian/pool/main/s/sgrep/sgrep_1.94a.orig.tar.gz"
	sgrep_cmd_tar := exec.Command("curl", "-L", sgrep_tar_url, "-o", "package.tar.gz")
	err := sgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sgrep_zip_url := "https://deb.debian.org/debian/pool/main/s/sgrep/sgrep_1.94a.orig.zip"
	sgrep_cmd_zip := exec.Command("curl", "-L", sgrep_zip_url, "-o", "package.zip")
	err = sgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sgrep_bin_url := "https://deb.debian.org/debian/pool/main/s/sgrep/sgrep_1.94a.orig.bin"
	sgrep_cmd_bin := exec.Command("curl", "-L", sgrep_bin_url, "-o", "binary.bin")
	err = sgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sgrep_src_url := "https://deb.debian.org/debian/pool/main/s/sgrep/sgrep_1.94a.orig.src.tar.gz"
	sgrep_cmd_src := exec.Command("curl", "-L", sgrep_src_url, "-o", "source.tar.gz")
	err = sgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sgrep_cmd_direct := exec.Command("./binary")
	err = sgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
