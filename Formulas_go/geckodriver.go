package main

// Geckodriver - WebDriver <-> Marionette proxy
// Homepage: https://github.com/mozilla/geckodriver

import (
	"fmt"
	
	"os/exec"
)

func installGeckodriver() {
	// Método 1: Descargar y extraer .tar.gz
	geckodriver_tar_url := "https://hg.mozilla.org/mozilla-central/archive/#{hg_revision}.zip/testing/geckodriver/"
	geckodriver_cmd_tar := exec.Command("curl", "-L", geckodriver_tar_url, "-o", "package.tar.gz")
	err := geckodriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geckodriver_zip_url := "https://hg.mozilla.org/mozilla-central/archive/#{hg_revision}.zip/testing/geckodriver/"
	geckodriver_cmd_zip := exec.Command("curl", "-L", geckodriver_zip_url, "-o", "package.zip")
	err = geckodriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geckodriver_bin_url := "https://hg.mozilla.org/mozilla-central/archive/#{hg_revision}.zip/testing/geckodriver/"
	geckodriver_cmd_bin := exec.Command("curl", "-L", geckodriver_bin_url, "-o", "binary.bin")
	err = geckodriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geckodriver_src_url := "https://hg.mozilla.org/mozilla-central/archive/#{hg_revision}.zip/testing/geckodriver/"
	geckodriver_cmd_src := exec.Command("curl", "-L", geckodriver_src_url, "-o", "source.tar.gz")
	err = geckodriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geckodriver_cmd_direct := exec.Command("./binary")
	err = geckodriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
