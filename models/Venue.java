/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package models;

import java.util.LinkedList;

/**
 *
 * @author johancanales
 */
public class Venue {
    private int idVenue;
    private int totalAsientos;
    private LinkedList<String> listaTipoAsiento;

    public int getIdVenue() {
        return idVenue;
    }

    public int getTotalAsientos() {
        return totalAsientos;
    }

    public LinkedList<String> getListaTipoAsiento() {
        return listaTipoAsiento;
    }

    public void setIdVenue(int idVenue) {
        this.idVenue = idVenue;
    }

    public void setTotalAsientos(int totalAsientos) {
        this.totalAsientos = totalAsientos;
    }

    public void setListaTipoAsiento(LinkedList<String> listaTipoAsiento) {
        this.listaTipoAsiento = listaTipoAsiento;
    }
}
