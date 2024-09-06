package main

// Keepassc - Curses-based password manager for KeePass v.1.x and KeePassX
// Homepage: https://github.com/raymontag/keepassc

import (
	"fmt"
	
	"os/exec"
)

func installKeepassc() {
	// Método 1: Descargar y extraer .tar.gz
	keepassc_tar_url := "https://files.pythonhosted.org/packages/c8/87/a7d40d4a884039e9c967fb2289aa2aefe7165110a425c4fb74ea758e9074/keepassc-1.8.2.tar.gz"
	keepassc_cmd_tar := exec.Command("curl", "-L", keepassc_tar_url, "-o", "package.tar.gz")
	err := keepassc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keepassc_zip_url := "https://files.pythonhosted.org/packages/c8/87/a7d40d4a884039e9c967fb2289aa2aefe7165110a425c4fb74ea758e9074/keepassc-1.8.2.zip"
	keepassc_cmd_zip := exec.Command("curl", "-L", keepassc_zip_url, "-o", "package.zip")
	err = keepassc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keepassc_bin_url := "https://files.pythonhosted.org/packages/c8/87/a7d40d4a884039e9c967fb2289aa2aefe7165110a425c4fb74ea758e9074/keepassc-1.8.2.bin"
	keepassc_cmd_bin := exec.Command("curl", "-L", keepassc_bin_url, "-o", "binary.bin")
	err = keepassc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keepassc_src_url := "https://files.pythonhosted.org/packages/c8/87/a7d40d4a884039e9c967fb2289aa2aefe7165110a425c4fb74ea758e9074/keepassc-1.8.2.src.tar.gz"
	keepassc_cmd_src := exec.Command("curl", "-L", keepassc_src_url, "-o", "source.tar.gz")
	err = keepassc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keepassc_cmd_direct := exec.Command("./binary")
	err = keepassc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
