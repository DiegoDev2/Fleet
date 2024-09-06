package main

// Sleuthkit - Forensic toolkit
// Homepage: https://www.sleuthkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installSleuthkit() {
	// Método 1: Descargar y extraer .tar.gz
	sleuthkit_tar_url := "https://github.com/sleuthkit/sleuthkit/releases/download/sleuthkit-4.12.1/sleuthkit-4.12.1.tar.gz"
	sleuthkit_cmd_tar := exec.Command("curl", "-L", sleuthkit_tar_url, "-o", "package.tar.gz")
	err := sleuthkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sleuthkit_zip_url := "https://github.com/sleuthkit/sleuthkit/releases/download/sleuthkit-4.12.1/sleuthkit-4.12.1.zip"
	sleuthkit_cmd_zip := exec.Command("curl", "-L", sleuthkit_zip_url, "-o", "package.zip")
	err = sleuthkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sleuthkit_bin_url := "https://github.com/sleuthkit/sleuthkit/releases/download/sleuthkit-4.12.1/sleuthkit-4.12.1.bin"
	sleuthkit_cmd_bin := exec.Command("curl", "-L", sleuthkit_bin_url, "-o", "binary.bin")
	err = sleuthkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sleuthkit_src_url := "https://github.com/sleuthkit/sleuthkit/releases/download/sleuthkit-4.12.1/sleuthkit-4.12.1.src.tar.gz"
	sleuthkit_cmd_src := exec.Command("curl", "-L", sleuthkit_src_url, "-o", "source.tar.gz")
	err = sleuthkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sleuthkit_cmd_direct := exec.Command("./binary")
	err = sleuthkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
	exec.Command("latte", "install", "ant").Run()
	fmt.Println("Instalando dependencia: afflib")
	exec.Command("latte", "install", "afflib").Run()
	fmt.Println("Instalando dependencia: libewf")
	exec.Command("latte", "install", "libewf").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
