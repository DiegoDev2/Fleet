package main

// Tarsnapper - Tarsnap wrapper which expires backups using a gfs-scheme
// Homepage: https://github.com/miracle2k/tarsnapper

import (
	"fmt"
	
	"os/exec"
)

func installTarsnapper() {
	// Método 1: Descargar y extraer .tar.gz
	tarsnapper_tar_url := "https://files.pythonhosted.org/packages/4e/c5/0a08950e5faba96e211715571c68ef64ee37b399ef4f0c4ab55e66c3c4fe/tarsnapper-0.5.0.tar.gz"
	tarsnapper_cmd_tar := exec.Command("curl", "-L", tarsnapper_tar_url, "-o", "package.tar.gz")
	err := tarsnapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tarsnapper_zip_url := "https://files.pythonhosted.org/packages/4e/c5/0a08950e5faba96e211715571c68ef64ee37b399ef4f0c4ab55e66c3c4fe/tarsnapper-0.5.0.zip"
	tarsnapper_cmd_zip := exec.Command("curl", "-L", tarsnapper_zip_url, "-o", "package.zip")
	err = tarsnapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tarsnapper_bin_url := "https://files.pythonhosted.org/packages/4e/c5/0a08950e5faba96e211715571c68ef64ee37b399ef4f0c4ab55e66c3c4fe/tarsnapper-0.5.0.bin"
	tarsnapper_cmd_bin := exec.Command("curl", "-L", tarsnapper_bin_url, "-o", "binary.bin")
	err = tarsnapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tarsnapper_src_url := "https://files.pythonhosted.org/packages/4e/c5/0a08950e5faba96e211715571c68ef64ee37b399ef4f0c4ab55e66c3c4fe/tarsnapper-0.5.0.src.tar.gz"
	tarsnapper_cmd_src := exec.Command("curl", "-L", tarsnapper_src_url, "-o", "source.tar.gz")
	err = tarsnapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tarsnapper_cmd_direct := exec.Command("./binary")
	err = tarsnapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: tarsnap")
exec.Command("latte", "install", "tarsnap")
}
