package main

// ErlangAT24 - Programming language for highly scalable real-time systems
// Homepage: https://www.erlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installErlangAT24() {
	// Método 1: Descargar y extraer .tar.gz
	erlangat24_tar_url := "https://github.com/erlang/otp/releases/download/OTP-24.3.4.17/otp_src_24.3.4.17.tar.gz"
	erlangat24_cmd_tar := exec.Command("curl", "-L", erlangat24_tar_url, "-o", "package.tar.gz")
	err := erlangat24_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erlangat24_zip_url := "https://github.com/erlang/otp/releases/download/OTP-24.3.4.17/otp_src_24.3.4.17.zip"
	erlangat24_cmd_zip := exec.Command("curl", "-L", erlangat24_zip_url, "-o", "package.zip")
	err = erlangat24_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erlangat24_bin_url := "https://github.com/erlang/otp/releases/download/OTP-24.3.4.17/otp_src_24.3.4.17.bin"
	erlangat24_cmd_bin := exec.Command("curl", "-L", erlangat24_bin_url, "-o", "binary.bin")
	err = erlangat24_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erlangat24_src_url := "https://github.com/erlang/otp/releases/download/OTP-24.3.4.17/otp_src_24.3.4.17.src.tar.gz"
	erlangat24_cmd_src := exec.Command("curl", "-L", erlangat24_src_url, "-o", "source.tar.gz")
	err = erlangat24_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erlangat24_cmd_direct := exec.Command("./binary")
	err = erlangat24_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
	fmt.Println("Instalando dependencia: wxwidgets")
exec.Command("latte", "install", "wxwidgets")
}
