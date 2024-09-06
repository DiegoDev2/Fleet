package main

// ZitaConvolver - Fast, partitioned convolution engine library
// Homepage: https://kokkinizita.linuxaudio.org/linuxaudio/

import (
	"fmt"
	
	"os/exec"
)

func installZitaConvolver() {
	// Método 1: Descargar y extraer .tar.gz
	zitaconvolver_tar_url := "https://kokkinizita.linuxaudio.org/linuxaudio/downloads/zita-convolver-4.0.3.tar.bz2"
	zitaconvolver_cmd_tar := exec.Command("curl", "-L", zitaconvolver_tar_url, "-o", "package.tar.gz")
	err := zitaconvolver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zitaconvolver_zip_url := "https://kokkinizita.linuxaudio.org/linuxaudio/downloads/zita-convolver-4.0.3.tar.bz2"
	zitaconvolver_cmd_zip := exec.Command("curl", "-L", zitaconvolver_zip_url, "-o", "package.zip")
	err = zitaconvolver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zitaconvolver_bin_url := "https://kokkinizita.linuxaudio.org/linuxaudio/downloads/zita-convolver-4.0.3.tar.bz2"
	zitaconvolver_cmd_bin := exec.Command("curl", "-L", zitaconvolver_bin_url, "-o", "binary.bin")
	err = zitaconvolver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zitaconvolver_src_url := "https://kokkinizita.linuxaudio.org/linuxaudio/downloads/zita-convolver-4.0.3.tar.bz2"
	zitaconvolver_cmd_src := exec.Command("curl", "-L", zitaconvolver_src_url, "-o", "source.tar.gz")
	err = zitaconvolver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zitaconvolver_cmd_direct := exec.Command("./binary")
	err = zitaconvolver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
}
